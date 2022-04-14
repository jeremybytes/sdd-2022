using PeopleViewer.Library;
using System.ComponentModel;

namespace PeopleViewer;

internal class MainWindowViewModel : INotifyPropertyChanged
{
    private IEnumerable<Person> cachedPeople = new List<Person>();

    public IEnumerable<Person> People { get; private set; }

    private bool dateFilterChecked;
    public bool DateFilterChecked
    {
        get => dateFilterChecked;
        set { dateFilterChecked = value; UpdateFilterAndSort(); }
    }

    private int dateFilterStartYear;
    public int DateFilterStartYear
    {
        get => dateFilterStartYear;
        set { dateFilterStartYear = value; UpdateFilterAndSort(); }
    }

    private int dateFilterEndYear;
    public int DateFilterEndYear
    {
        get => dateFilterEndYear;
        set { dateFilterEndYear = value; UpdateFilterAndSort(); }
    }

    private bool nameFilterChecked;
    public bool NameFilterChecked
    {
        get => nameFilterChecked;
        set { nameFilterChecked = value; UpdateFilterAndSort(); }
    }

    private string nameFilterValue;
    public string NameFilterValue
    {
        get => nameFilterValue;
        set { nameFilterValue = value; UpdateFilterAndSort(); }
    }

    private bool familyNameSortChecked;
    public bool FamilyNameSortChecked
    {
        get => familyNameSortChecked;
        set
        {
            UncheckAllSort();
            familyNameSortChecked = value;
            UpdateFilterAndSort();
        }
    }

    private bool givenNameSortChecked;
    public bool GivenNameSortChecked
    {
        get => givenNameSortChecked;
        set
        {
            UncheckAllSort();
            givenNameSortChecked = value;
            UpdateFilterAndSort();
        }
    }

    private bool startDateSortChecked;
    public bool StartDateSortChecked
    {
        get => startDateSortChecked;
        set
        {
            UncheckAllSort();
            startDateSortChecked = value;
            UpdateFilterAndSort();
        }
    }

    private bool ratingSortChecked;
    public bool RatingSortChecked
    {
        get => ratingSortChecked;
        set
        {
            UncheckAllSort();
            ratingSortChecked = value;
            UpdateFilterAndSort();
        }
    }

    public MainWindowViewModel()
    {
        dateFilterStartYear = 1985;
        dateFilterEndYear = 2000;
        nameFilterValue = "John";
        People = cachedPeople;
    }

    public void RefreshData()
    {
        var reader = new PeopleReader();
        reader.GetPeopleCompleted += (s, e) =>
        {
            cachedPeople = e.Result;
            ResetFilters();
            UpdateFilterAndSort();
        };
        reader.GetPeopleAsync();
    }

    private void ResetFilters()
    {
        UncheckAllSort();
        dateFilterChecked = false;
        nameFilterChecked = false;
        UpdateFilterAndSort();
    }

    private void UncheckAllSort()
    {
        familyNameSortChecked = false;
        givenNameSortChecked = false;
        startDateSortChecked = false;
        ratingSortChecked = false;
    }

    private void UpdateFilterAndSort()
    {
        People = cachedPeople;

        if (DateFilterChecked)
            People = People
                .Where(p => p.StartDate.Year >= dateFilterStartYear)
                .Where(p => p.StartDate.Year <= dateFilterEndYear);

        if (nameFilterChecked)
            People = People.Where(p => p.GivenName == NameFilterValue);

        if (FamilyNameSortChecked)
            People = People.OrderBy(p => p.FamilyName);

        if (GivenNameSortChecked)
            People = People.OrderBy(p => p.GivenName).ThenBy(p => p.FamilyName);

        if (StartDateSortChecked)
            People = People.OrderBy(p => p.StartDate);

        if (RatingSortChecked)
            People = People.OrderByDescending(p => p.Rating);

        RaisePropertyChanged();
    }

    #region INotifyPropertyChanged Members

    public event PropertyChangedEventHandler? PropertyChanged;
    private void RaisePropertyChanged(string? propertyName = null)
    {
        PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
    }

    #endregion
}
